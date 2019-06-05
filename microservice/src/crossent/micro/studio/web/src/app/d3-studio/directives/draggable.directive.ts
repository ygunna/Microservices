import { Directive, Input, Output, EventEmitter, ElementRef, OnInit, HostListener } from '@angular/core';
import { D3Service } from '../shared/d3-studio.service';
import { Link } from '../shared/link.model';
import { EditComponent } from '../../compose/edit/edit.component';

@Directive({
  selector: '[draggableNode]'
})
export class DraggableDirective implements OnInit {
  @Input('draggableNode') draggableNode: ElementRef;
  @Input('droppedNodes') droppedNodes: Node[];
  @Input('linkPaths') linkPaths: Link[];
  @Output() clickRemove = new EventEmitter();

  constructor(private d3Service: D3Service, private _element: ElementRef, private editComponent: EditComponent) { }

  @HostListener('document:keypress', ['$event'])
  handleKeyboardEvent(event: KeyboardEvent) {
    if(event.key === 'Delete') {
      if(D3Service.selectedElement != null) {
        if(D3Service.selectedElement.type == 'line') {
          this.linkPaths.forEach(link => {
            if(link.id == D3Service.selectedElement.element.id && link.type == 'Service') {
              this.linkPaths.splice(this.linkPaths.indexOf(link), 1);
              return;
            }
          });
        } //else if(D3Service.selectedElement.type == 'node') {
          //this.clickRemove.emit(D3Service.selectedElement.element);
        //}
      }
    }
  }

  ngOnInit() {
    this.d3Service.applyDraggableBehaviour(this._element.nativeElement, this.draggableNode, this.droppedNodes, this.linkPaths, this.editComponent);
  }
}
