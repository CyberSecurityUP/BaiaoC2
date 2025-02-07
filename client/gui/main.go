package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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

// Start inicia a interface gráfica do BaiaoC2
func Start() {
	// Inicializar o aplicativo
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("BaiaoC2 - Command & Control")

	// Tabela de Listeners
	listenerTable := widget.NewTable(
		func() (int, int) { return len(listeners) + 1, 6 }, // +1 para cabeçalhos
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			if id.Row == 0 {
				// Cabeçalhos da tabela
				headers := []string{"Name", "Protocol", "Host", "PortBind", "PortConn", "Status"}
				label.SetText(headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				// Dados dos listeners
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

	// Funções dos Botões
	addButton := widget.NewButton("Add", func() { showAddDialog(myWindow, listenerTable) })
	removeButton := widget.NewButton("Remove", func() { removeListener(myWindow, listenerTable) })

	// Logs no painel lateral
	eventLogs := widget.NewMultiLineEntry()
	eventLogs.SetText("[+] Event Viewer Logs\n")
	eventLogs.Disable()

	// Rodapé (Status Bar)
	statusBar := widget.NewLabel("Status: Connected to Teamserver")

	// Layout Principal
	buttons := container.NewHBox(addButton, removeButton)
	mainContent := container.NewBorder(
		container.NewVBox(listenerTable, buttons), // Parte central
		statusBar,                                // Rodapé
		nil,                                      // Sem menu lateral esquerdo
		eventLogs,                                // Logs à direita
		nil,                                      // Sem cabeçalho adicional
	)

	// Configurar janela principal
	myWindow.SetContent(mainContent)
	myWindow.Resize(fyne.NewSize(1280, 720))
	myWindow.ShowAndRun()
}

// showAddDialog exibe o diálogo para adicionar listeners
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
			dialog.ShowInformation("Success", "Listener added successfully!", window)
		}
	}, window)
}

// removeListener remove o listener selecionado
func removeListener(window fyne.Window, table *widget.Table) {
	if len(listeners) == 0 {
		dialog.ShowInformation("Error", "No listeners to remove.", window)
		return
	}

	// Remove o último listener por simplicidade
	listeners = listeners[:len(listeners)-1]
	table.Refresh()
	dialog.ShowInformation("Success", "Listener removed successfully!", window)
}
