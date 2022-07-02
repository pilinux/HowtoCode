import { 
  AfterContentChecked, 
  AfterContentInit, 
  AfterViewChecked, 
  AfterViewInit, 
  Component, 
  ContentChild, 
  DoCheck, 
  ElementRef, 
  Input, 
  OnChanges, 
  OnDestroy, 
  OnInit, 
  SimpleChanges, 
  ViewChild
} from '@angular/core';

@Component({
  selector: 'app-server-element',
  templateUrl: './server-element.component.html',
  styleUrls: ['./server-element.component.css']
})
export class ServerElementComponent implements 
OnInit, 
OnChanges, 
DoCheck, 
AfterContentInit, 
AfterContentChecked,
AfterViewInit,
AfterViewChecked,
OnDestroy {
  // property to access from outside of this component
  @Input('srvElement') element: {type: string, name: string, content: string};
  @Input() name: string;

  @ViewChild('heading') header: ElementRef;
  @ContentChild('contentParagraph') paragraph: ElementRef;

  constructor() {
    console.log('constructor called!');
   }

   ngOnChanges(changes: SimpleChanges): void {
    console.log('ngOnChanges called!');
    console.log(changes);
   }

  ngOnInit(): void {
    console.log('ngOnInit called!');
    // check when @ViewChild can access the template
    // console.log('textContent:' + this.header.nativeElement.textContent);
    // will throw error

    // check when @ContentChild can access the template
    // console.log('textContent of paragraph:' + this.paragraph.nativeElement.textContent);
    // will throw error
  }

  ngDoCheck(): void {
    console.log('ngDoCheck called!');
  }

  ngAfterContentInit(): void {
    console.log('ngAfterContentInit called!');
    // check when @ContentChild can access the template
    console.log('textContent of paragraph:' + this.paragraph.nativeElement.textContent);
  }

  ngAfterContentChecked(): void {
    console.log('ngAfterContentChecked called!');
  }

  ngAfterViewInit(): void {
    console.log('ngAfterViewInit called!');
    // check when @ViewChild can access the template
    console.log('textContent:' + this.header.nativeElement.textContent);
  }

  ngAfterViewChecked(): void {
    console.log('ngAfterViewChecked called!');
  }

  ngOnDestroy(): void {
    console.log('ngOnDestroy called!');
  }

}
