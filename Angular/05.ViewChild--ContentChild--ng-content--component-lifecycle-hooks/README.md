# @ViewChild, @ContentChild, ng-content, component lifecycle hooks

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 14.0.2.

## File structure: previous project

_Note:_ Continuation of `04.split.app.into.components`

```
src
└───app
    └---app.component.css
    └---app.component.html
    └---app.component.ts
    └---app.module.ts
    │
    └───cockpit
    │   └---cockpit.component.css
    │   └---cockpit.component.html
    │   └---cockpit.component.ts
    │
    └───server-element
        └---server-element.component.css
        └---server-element.component.html
        └---server-element.component.ts
```

## Local references in templates

In `cockpit.component.html`, _two-way data-binding_ is used to get
_server name_ and _server content_. Since only at the point of clicking
the button, the data is used or saved, it would be enough to get
the value of the input at this point of time by removing
_two-way data-binding_ and placing a _local reference_ on that element.

A _local reference_ can be placed on any HTML element. It only
holds a reference to that HTML element, but not the actual value.

**cockpit.component.html:**

```
<p>Add new Servers or blueprints!</p>
<label>Server Name</label>
<!--input type="text" class="form-control" [(ngModel)]="newServerName"-->
<!--using local reference-->
<!--this reference will hold a reference to this element, not the actual value-->
<input type="text" class="form-control" #serverNameInput>
<label>Server Content</label>
<input type="text" class="form-control" [(ngModel)]="newServerContent">
<br>
<button class="btn btn-primary" (click)="onAddServer(serverNameInput)">Add Server</button>&nbsp;
<button class="btn btn-primary" (click)="onAddBlueprint(serverNameInput)">Add Server Blueprint</button>
```

**cockpit.component.ts:**

```
import { Component, EventEmitter, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-cockpit',
  templateUrl: './cockpit.component.html',
  styleUrls: ['./cockpit.component.css']
})
export class CockpitComponent implements OnInit {
  // newServerName = '';
  newServerContent = '';

  // properties that emit events
  // EventEmitter is an object in angular framework that allows to emit events
  // @Output(): listenable from outside the component
  @Output('srvCreated') serverCreated = new EventEmitter<{serverName: string, serverContent: string}>();
  @Output('bpCreated') blueprintCreated = new EventEmitter<{serverName: string, serverContent: string}>();

  constructor() { }

  ngOnInit(): void {
  }

  onAddServer(nameInput: HTMLInputElement) {
    // console.log(nameInput);
    // console.log(nameInput.value);
    this.serverCreated.emit({
      serverName: nameInput.value,
      serverContent: this.newServerContent
    });
  }

  onAddBlueprint(nameInput: HTMLInputElement) {
    this.blueprintCreated.emit({
      serverName: nameInput.value,
      serverContent: this.newServerContent
    });
  }
}
```

## @ViewChild: access the template and DOM

**cockpit.component.html:**

```
<p>Add new Servers or blueprints!</p>
<label>Server Name</label>
<!--input type="text" class="form-control" [(ngModel)]="newServerName"-->
<!--using local reference-->
<!--this reference will hold a reference to this element, not the actual value-->
<input type="text" class="form-control" #serverNameInput>
<label>Server Content</label>
<!--input type="text" class="form-control" [(ngModel)]="newServerContent"-->
<input type="text" class="form-control" #serverContentInput>
<br>
<button class="btn btn-primary" (click)="onAddServer(serverNameInput)">Add Server</button>&nbsp;
<button class="btn btn-primary" (click)="onAddBlueprint(serverContentInput)">Add Server Blueprint</button>
```

**cockpit.component.ts:**

```
import { Component, ElementRef, EventEmitter, OnInit, Output, ViewChild } from '@angular/core';

@Component({
  selector: 'app-cockpit',
  templateUrl: './cockpit.component.html',
  styleUrls: ['./cockpit.component.css']
})
export class CockpitComponent implements OnInit {
  // newServerName = '';
  // newServerContent = '';
  @ViewChild('serverContentInput') serverContentInput: ElementRef;

  // properties that emit events
  // EventEmitter is an object in angular framework that allows to emit events
  // @Output(): listenable from outside the component
  @Output('srvCreated') serverCreated = new EventEmitter<{serverName: string, serverContent: string}>();
  @Output('bpCreated') blueprintCreated = new EventEmitter<{serverName: string, serverContent: string}>();

  constructor() { }

  ngOnInit(): void {
  }

  onAddServer(nameInput: HTMLInputElement) {
    // console.log(nameInput);
    // console.log(nameInput.value);
    // console.log(this.serverContentInput);
    this.serverCreated.emit({
      serverName: nameInput.value,
      serverContent: this.serverContentInput.nativeElement.value
    });
  }

  onAddBlueprint(nameInput: HTMLInputElement) {
    this.blueprintCreated.emit({
      serverName: nameInput.value,
      serverContent: this.serverContentInput.nativeElement.value
    });
  }
}
```

## ng-content: projecting content into components

**server-element.component.html:**

```
<div class="panel panel-default">
    <div class="panel-heading">{{ element.name }}</div>
    <div class="panel-body">
        <!--special directive-->
        <ng-content></ng-content>
    </div>
</div>
```

**app.component.html:**

```
<div class="container">

  <div class="row">
    <div class="col-xs-12">
      <app-cockpit (srvCreated)="onServerAdded($event)" (bpCreated)="onBlueprintAdded($event)">
      </app-cockpit>
    </div>
  </div>
  <hr>

  <div class="row">
    <div class="col-xs-12">
      <app-server-element *ngFor="let serverElement of serverElements" [srvElement]="serverElement">
        <p>
          <strong *ngIf="serverElement.type === 'server'" style="color: red">{{ serverElement.content }}</strong>
          <em *ngIf="serverElement.type === 'blueprint'">{{ serverElement.content }}</em>
        </p>
      </app-server-element>
    </div>
  </div>
</div>
```

## Component lifecycle hooks

[Official doc][01]

- `ngOnChanges`: called after a bound input property changes
- `ngOnInit`: called once the component is initialized
- `ngDoCheck`: called during every change detection run
- `ngAfterContentInit`: called after content (ng-content) has been projected into view
- `ngAfterContentChecked`: called every time the projected content has been checked
- `ngAfterViewInit`: called after the component’s view (and child views) has been initialized
- `ngAfterViewChecked`: called every time the view (and child views) have been checked
- `ngOnDestroy`: called once the component is about to be destroyed

The following files contain lifecycle hooks examples:

```
src
└───app
    └---app.component.html
    └---app.component.ts
    │
    └───server-element
        └---server-element.component.html
        └---server-element.component.ts
```

[01]: https://angular.io/guide/lifecycle-hooks
