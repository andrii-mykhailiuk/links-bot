package telegram

const msgHelp = `I can save and keep your pages. Also I can offer you them to read.

In order to save the page, just send me a link to it.

In order to get a random page fro your list, send me command /rnd.
Caution! After that, this page will be removed from your list!
`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "You don't have saved pages"
	msgSaved          = "Saved!"
	msgAlreadyExists  = "You have already had this page in your list."
)
