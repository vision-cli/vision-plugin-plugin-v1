# ![logo](./images/vision-logo.svg "Vision") &nbsp; Vision Plugin - Plugin

This plugin creates a standard plugin template

Vision plugins require golang (https://go.dev) to be installed

## Install

Install the plugin with

```
go install github.com/vision-cli/vision-plugin-plugin-v1@latest
```

You will now see the plugin-plugin help in the vision cli

```
vision plugin --help
```

You will now see the plugin available in vision's help

```
vision --help
```

The plugin will be listed in plugins list

```
vision plugins list
```

## Run

Create a plugin project

```
vision init <plugin project name>
```

Switch to the plugin project directory

```
cd <plugin project name>
```

Initialise the plugin plugin

```
vision plugin init <plugin module name e.g. github.com/vision-cli/vision-plugin-plugin-v1>
```

Generate the code

```
vision plugin generate
```

This will generate the plugin template code in the current folder. Follow the instructions in the README.md
