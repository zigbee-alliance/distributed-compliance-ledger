# Distributed Compliance Ledger UI

__This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 8.3.20.__

- Run `ng serve` for a dev server
- Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.
- Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).
- Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

__Read more about the project structure:__

- https://medium.com/@motcowley/angular-folder-structure-d1809be95542
- https://itnext.io/choosing-a-highly-scalable-folder-structure-in-angular-d987de65ec7

__Dev environment__

You will want nginx to add 'allow cross origin' headers to `dclcli` responses. Find nginx configuration template in `templates` and place it into `/etc/nginx/conf.d`.

__Deployment__

- Read more about deployment in `ansible/README.md`.

__Configuration__

Configuration is located in: `src/environments/environment.tx` and `src/environments/environment.prod.tx` for dev and prod environments correspondingly.
There are two important settings here:

- `apiUrl`: base path of the running `dclcli rest-server`;
- `chainId`: chain id, `dclchain` by default.
