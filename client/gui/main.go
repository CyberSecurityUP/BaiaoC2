package gui

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

var _ = json.Marshal // Isso engana o compilador dizendo que `json` está sendo usado

// Listener representa um listener no BaiaoC2
type Listener struct {
	Name       string
	Protocol   string
	Host       string
	PortBind   string
	PortConn   string
	Status     string
}

// Lista de listeners (inicia vazia)
var listeners = []Listener{}

// **Start inicia a interface gráfica do BaiaoC2**
func Start() {
	// Inicializar o aplicativo
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("BaiaoC2 - Command & Control")

	// **Inicializa os logs**
	logs := widget.NewMultiLineEntry()
	logs.Disable()
	logs.SetText("[+] Event Viewer Logs\n")

	// **Abas principais**
	tabs := container.NewAppTabs()

	// **Aba de Listeners**
	listenerTable := createListenerTable(myWindow) // Corrigido: Passa `myWindow`, não `logs`
	tabs.Append(container.NewTabItem("Listeners", listenerTable))

	// **Aba de Session View**
	sessionView := widget.NewLabel("Session View Content")
	tabs.Append(container.NewTabItem("Session View", sessionView))

	// **Aba de Teamserver Chat**
	teamserverChat := widget.NewMultiLineEntry()
	teamserverChat.SetText("Teamserver Chat Content")
	teamserverChat.Disable()
	tabs.Append(container.NewTabItem("Teamserver Chat", container.NewScroll(teamserverChat)))

	// **Aba de Loot**
	loot := widget.NewLabel("Loot Content")
	tabs.Append(container.NewTabItem("Loot", loot))

	// **Aba de Event Viewer**
	tabs.Append(container.NewTabItem("Event Viewer", container.NewScroll(logs)))

	// **Menu no topo com submenus corrigidos**
	menu := fyne.NewMainMenu(
		fyne.NewMenu("View",
			fyne.NewMenuItem("Listeners", func() { tabs.SelectTabIndex(0) }),
			fyne.NewMenuItem("Session View", func() { tabs.SelectTabIndex(1) }),
			fyne.NewMenuItem("Teamserver Chat", func() { tabs.SelectTabIndex(2) }),
			fyne.NewMenuItem("Loot", func() { tabs.SelectTabIndex(3) }),
			fyne.NewMenuItem("Event Viewer", func() { tabs.SelectTabIndex(4) }),
		),
		fyne.NewMenu("Attack",
			fyne.NewMenuItem("Generate Payload", func() { showMessage(myWindow, "Generate Payload feature coming soon!") }),
			fyne.NewMenuItem("Upload Payload", func() { showMessage(myWindow, "Upload Payload feature coming soon!") }),
			fyne.NewMenuItem("Execute Task", func() { showMessage(myWindow, "Execute Task feature coming soon!") }),
		),
		fyne.NewMenu("Scripts",
			fyne.NewMenuItem("Manage Scripts", func() { showMessage(myWindow, "Manage Scripts feature coming soon!") }),
			fyne.NewMenuItem("Run Script", func() { showMessage(myWindow, "Run Script feature coming soon!") }),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("Documentation", func() { showMessage(myWindow, "Documentation coming soon!") }),
			fyne.NewMenuItem("About", func() {
				dialog.ShowInformation("About BaiaoC2", "BaiaoC2 - Version 1.0\nDeveloped for advanced C2 operations.", myWindow)
			}),
		),
	)
	myWindow.SetMainMenu(menu)

	// **Configurar janela principal**
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(1280, 720))
	myWindow.ShowAndRun()
}

// **Cria a tabela de listeners**
func createListenerTable(window fyne.Window) *fyne.Container {
	listenerTable := widget.NewTable(
		func() (int, int) { return len(listeners) + 1, 6 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				headers := []string{"Name", "Protocol", "Host", "PortBind", "PortConn", "Status"}
				label.SetText(headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				listener := listeners[id.Row-1]
				switch id.Col {
				case 0:
					label.SetText(listener.Name)
				case 1:
					label.SetText(listener.Protocol)
				case 2:
					label.SetText(listener.Host)
				case 3:
					label.SetText(listener.PortBind)
				case 4:
					label.SetText(listener.PortConn)
				case 5:
					label.SetText(listener.Status)
				}
			}
		},
	)

	listenerTable.SetColumnWidth(0, 150)
	listenerTable.SetColumnWidth(1, 120)
	listenerTable.SetColumnWidth(2, 200)
	listenerTable.SetColumnWidth(3, 100)
	listenerTable.SetColumnWidth(4, 100)
	listenerTable.SetColumnWidth(5, 80)

	// **Botão para adicionar listeners**
	addButton := widget.NewButton("Add Listener", func() {
		showAddDialog(window, listenerTable) // Certifique-se de passar `window`
	})

	return container.NewBorder(nil, addButton, nil, nil, listenerTable)
}

// **Exibe o diálogo para adicionar listeners**
func showAddDialog(window fyne.Window, table *widget.Table) {
	nameEntry := widget.NewEntry()
	protocolEntry := widget.NewEntry()
	hostEntry := widget.NewEntry()
	portBindEntry := widget.NewEntry()
	portConnEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Protocol", protocolEntry),
		widget.NewFormItem("Host", hostEntry),
		widget.NewFormItem("PortBind", portBindEntry),
		widget.NewFormItem("PortConn", portConnEntry),
	)

	// **Evita erro caso a janela seja `nil`**
	if window == nil {
		log.Println("[!] Erro: A janela principal não foi inicializada corretamente.")
		return
	}

	dialog.ShowCustomConfirm("Add Listener", "Add", "Cancel", form, func(confirmed bool) {
		if confirmed {
			newListener := Listener{
				Name:     nameEntry.Text,
				Protocol: protocolEntry.Text,
				Host:     hostEntry.Text,
				PortBind: portBindEntry.Text,
				PortConn: portConnEntry.Text,
				Status:   "Online",
			}
			listeners = append(listeners, newListener)
			table.Refresh()
		}
	}, window) // Certifique-se de passar a `window` corretamente
}

// **Exibe mensagens temporárias**
func showMessage(window fyne.Window, message string) {
	dialog.ShowInformation("Info", message, window)
}
