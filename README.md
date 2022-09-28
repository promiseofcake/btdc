# btdc

A macOS application to unpair all but an allow-list of bluetooth devices.

Useful in shared-office situations where you use different input devices on subsequent days.


## Requirements

MacOS and blueutil (https://formulae.brew.sh/formula/blueutil) installed: 

```bash
brew install blueutil
````

## Usage

For now, just compile it in the repo and create an `allow.yml` with the following format:

```yaml
devices:
- 04-fe-a1-4e-22-51
- 0c-d7-46-ee-22-49
- 28-11-a5-44-22-47
- 38-ec-0d-60-22-46
- 4c-57-ca-ed-22-45
- 74-65-0c-b4-22-44
```

Then run the program:

```bash
go run . 
Successfully unpaired Tom Nook's Mouse [0c-e4-41-00-50-99]!
Successfully unpaired Isabelle's Keyboard [64-0b-d7-1f-02-99]!
Unpaired 2 devices, failed to unpair 0, devices!
```

## Follow-up

* Improve tool to allow interactive creation of the allow.yaml file based upon paired devices.
* Reduce reliance on third party tools (blueutil)
