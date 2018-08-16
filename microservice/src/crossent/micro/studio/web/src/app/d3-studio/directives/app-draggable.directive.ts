import { Directive, ElementRef, HostListener, Input } from '@angular/core';

@Directive({
  selector: '[appDraggable]'
})
export class AppDraggableDirective {
  @Input('appDraggable') appDraggable: ElementRef;

  constructor(private _element: ElementRef) {
    this._element.nativeElement.setAttribute('draggable', true);
  }

  @HostListener('dragstart', ['$event'])
  onDragStart(event) {
    event.dataTransfer.setData('application/text', event.target.id);
    this.appDraggable['shape'] = this._element.nativeElement.dataset.shape;
    this.appDraggable['color'] = this._element.nativeElement.dataset.color;
    this.appDraggable['type'] = this._element.nativeElement.dataset.type;
    event.dataTransfer.setData("nodeData", JSON.stringify(this.appDraggable));
    // var img = new Image();
    // img.src = 'https://static.wixstatic.com/media/2cd43b_fe13fdc4a31d4d94892244d99b122175~mv2.png/v1/fill/w_189,h_189,al_c,usm_0.66_1.00_0.01/2cd43b_fe13fdc4a31d4d94892244d99b122175~mv2.png';
    // event.dataTransfer.setDragImage(img, 0, 0);
  }

  @HostListener('document:dragover', ['$event'])
  onDragOver(event) {
    event.preventDefault();
  }
}
