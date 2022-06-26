# CmpDatabindingStart

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 14.0.2.

## Getting started

`ng new cmp-databinding-start --no-strict`

- NO angular routing
- Stylesheet format: CSS

`cd cmp-databinding-start`

`npm install --save bootstrap@3`

angular.json:

```
"styles": [
    "node_modules/bootstrap/dist/css/bootstrap.min.css",
    "src/styles.css"
],
```

```
cmp-databinding-start
└───src
│   └───app
│       └---app.component.css
│       └---app.component.html
│       └---app.component.ts
│       └---app.module.ts
```

app.component.css:

```
.container {
    margin-top: 30px;
  }

  p {
    color: blue;
  }
```

app.component.html:

```
<div class="container">
  <div class="row">
    <div class="col-xs-12">
      <p>Add new Servers or blueprints!</p>
      <label>Server Name</label>
      <input type="text" class="form-control" [(ngModel)]="newServerName">
      <label>Server Content</label>
      <input type="text" class="form-control" [(ngModel)]="newServerContent">
      <br>
      <button class="btn btn-primary" (click)="onAddServer()">Add Server</button>&nbsp;
      <button class="btn btn-primary" (click)="onAddBlueprint()">Add Server Blueprint</button>
    </div>
  </div>
  <hr>
  <div class="row">
    <div class="col-xs-12">
      <div class="panel panel-default" *ngFor="let element of serverElements">
        <div class="panel-heading">{{ element.name }}</div>
        <div class="panel-body">
          <p>
            <strong *ngIf="element.type === 'server'" style="color: red">{{ element.content }}</strong>
            <em *ngIf="element.type === 'blueprint'">{{ element.content }}</em>
          </p>
        </div>
      </div>
    </div>
  </div>
</div>
```

app.component.ts:

```
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'cmp-databinding-start';

  serverElements = [];
  newServerName = '';
  newServerContent = '';

  onAddServer() {
    this.serverElements.push({
      type: 'server',
      name: this.newServerName,
      content: this.newServerContent
    });
  }

  onAddBlueprint() {
    this.serverElements.push({
      type: 'blueprint',
      name: this.newServerName,
      content: this.newServerContent
    });
  }
}
```

app.module.ts:

```
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
```

## Custom properties and custom events binding

- Javascript object is noted by `{}`, example: `element: {}`
- By default, all properties are accessible only from inside the component.
- With decorator, for example `@Input()`, property can be exposed globally.
  Example: ` @Input() element: {type: string, name: string, content: string};`.
  Here, `element` is accessible from any component.
- Alias can be assigned to custom properties.
  Example: ` @Input(srvElement) element: {type: string, name: string, content: string};`.
  Here, `element` is accessible from the component itself. To access it from other
  components, `srvElement` alias has to be used.

## Split app into components

`ng g c cockpit --skip-tests`

`ng g c server-element --skip-tests`

```
cmp-databinding-start
└───src
│   └───app
│       └---app.component.css
│       └---app.component.html
│       └---app.component.ts
│       └---app.module.ts
│       │
│       └───cockpit
│       │   └---cockpit.component.css
│       │   └---cockpit.component.html
│       │   └---cockpit.component.ts
│       │
│       └───server-element
│           └---server-element.component.css
│           └---server-element.component.html
│           └---server-element.component.ts
```

_Note:_ Final codes are in the `src` directory.

## View encapsulation

- Not a default CSS behavior
- Angular, by default, encapsulates; CSS belongs only to the component it is defined

```
import { ... ..., ViewEncapsulation } from '@angular/core';

@Component({
 ...,
 encapsulation: ViewEncapsulation.Emulated
})
```

- Default: `encapsulation: ViewEncapsulation.Emulated`
- Make CSS property global: `encapsulation: ViewEncapsulation.None`
- Not all browsers support emulated view: `encapsulation: ViewEncapsulation.ShadowDom`

## Final result

![App][01]

[01]: final-app.png
