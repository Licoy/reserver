[中文](./README.md) | English

## :memo: Introduce

`reserver` It is a local server with real-time reload function designed for static website preview or development.

## :tada: Use

Go to [Releases](https://github.com/Licoy/reserver/releases) to select the latest release version of your operating
system, how to use it:

- Place this application in the root directory of your project.
- Add this program to the global, open the terminal in your project and start `reserver [options]`.

## :wrench: Options

```text
  -p, --port       listen port (default: 8080)
  -r, --root       root directory
  -H, --host       bind host address (default: 0.0.0.0)
      --no-browser  don't auto opening browser
      --no-watch    don't listen for file changes
      --browser    specify the browser you want to use
  -P, --path       default open link path
      --hide-log    displays the change log of the observation path
  -w, --wait       wait for the specified time before reloading (default: 100ms)
  -i, --ignore     multiple observation paths are allowed to be ignored, ex:
                    -i /a -i /b
  -v, --version     view current version
  -h, --help        Show this help message
```

## :label: Feedback
- If you encounter problems during use, you can directly submit [issue](https://github.com/Licoy/reserver/issues/new) for discussion or feedback.


## :page_facing_up: License
[MIT](./LICENSE)