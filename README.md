# ServerFan GitHub Action
This action can send notify or command to [server.fan](https://server.fan)
service in your GitHub action.

Learn more about [server.fan](https://server.fan/docs)

**Notice**: The server.fan service is available only in China now, it will release a global version later.

## Inputs

- token: **Required**, Your User-Token of server.fan, please use GitHub Secrets
- success: The message text you want see if success
- failed: The message text you want see if failed
- command: You can trigger a command after notification, such as a CD process, only run in success.
- source: The source text you want see, default is the repo name

## Example

```yaml
- name: Notification
  uses: hack-fan/skadi-action@v1
  with:
    token: ${{ secrets.SKADI_TOKEN }}
    success: 'Success'
    failed: 'Failed'
```
