import { Component, OnInit } from '@angular/core';
import { Recipe } from '../recipe.model';

@Component({
  selector: 'app-recipe-list',
  templateUrl: './recipe-list.component.html',
  styleUrls: ['./recipe-list.component.css']
})
export class RecipeListComponent implements OnInit {
  recipes: Recipe[] = [
    new Recipe('A Test Recipe', 'This is simply a test', 'https://cdn.pilinux.workers.dev/images/GoREST/logo/GoREST-Logo.png'),
    new Recipe('A second Test Recipe', 'This is simply another test', 'https://cdn.pilinux.workers.dev/images/GoREST/logo/GoREST-Logo.png')
  ];

  constructor() { }

  ngOnInit(): void {
  }

}
