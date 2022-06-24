# RecipeApp

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 14.0.2.

## Getting started

`ng new recipe-app --no-strict`

- NO angular routing
- Stylesheet format: CSS

`cd recipe-app`

`npm install --save bootstrap@3`

app.module.ts:

```
import { FormsModule } from '@angular/forms';

@NgModule({
      imports: [
        ...,
        FormsModule
  ],
})
```

angular.json:

```
"styles": [
    "node_modules/bootstrap/dist/css/bootstrap.min.css",
    "src/styles.css"
],
```

`ng g c header -skip-tests`

`ng g c recipes -skip-tests`

`ng g c recipes/recipe-list -skip-tests`

`ng g c recipes/recipe-detail -skip-tests`

`ng g c recipes/recipe-list/recipe-item -skip-tests`

`ng g c shopping-list -skip-tests`

`ng g c shopping-list/shopping-edit -skip-tests`

Replace files from this repo:

```
recipe-app
└───src
│   └───app
│       └---app.component.html
│       │
│       └───header
│       │   └---header.component.html
│       │   └---header.component.ts
│       │
│       └───recipes
│       │   └---recipe.model.ts
│       │   └---recipes.component.html
│       │   │
│       │   └───recipe-detail
│       │   │   └---recipe-detail.component.html
│       │   │
│       │   └───recipe-list
│       │       └---recipe-list.component.html
│       │       └---recipe-list.component.ts
│       │
│       └───shared
│       │   └---ingredient.model.ts
│       │
│       └───shopping-list
│           └---shopping-list.component.html
│           └---shopping-list.component.ts
│           │
│           └───shopping-edit
│               └---shopping-edit.component.html
```

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The application will automatically reload if you change any of the source files.

## Final result

![App][01]

[01]: final-app.png
