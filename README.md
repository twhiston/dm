# DM
Tom Whiston <tom.whiston@gmail.com>

## Bootstrapper for Docker for Mac

## Adding new commands

All commands take basically the same form if they need to interact with a process, and use listeners to trigger actions in the process types
Take this example from the start command

```
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Docker 4 Mac local Environment",
	Long: `Start the local environment by running all of the Extensions`,
	Run: func(cmd *cobra.Command, args []string) {
		o := SetUpListeners() //All commands should call this to add all listeners
		strict, _ := cmd.PersistentFlags().GetBool("strict") //if you need to check requirements you should pass in the strict flag
		o.Trigger("check-requirements", strict) //Trigger a requirements check
		o.Trigger("start", cfgFilePath) //Trigger the actual command
	},
}
```

## Existing Callbacks for listeners and arguments

'check-requirements', strict bool
'start', cfgFilePath string
'stop'
'init'