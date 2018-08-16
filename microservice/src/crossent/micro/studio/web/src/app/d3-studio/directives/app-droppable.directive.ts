import { Directive, HostListener, Input, Output, EventEmitter } from '@angular/core';
import { D3Service } from '../shared/d3-studio.service';
import { Node } from '../shared/node.model';

@Directive({
  selector: '[appDroppable]'
})
export class AppDroppableDirective {
  // @Output() appDroppable = new EventEmitter();
  @Input('droppedNodes') droppedNodes: Node;
  @Input('msaName') msaName: string;

  constructor(private d3Service: D3Service) {}

  @HostListener('drop', ['$event'])
  onDrop(event) {
    const nodeData = event.dataTransfer.getData('nodeData');
    var jsonObject : any = JSON.parse(nodeData);
    this.d3Service.applyAppDroppableBehaviour(event, jsonObject, this.droppedNodes, this.msaName);
  }

  @HostListener('mousedown', ['$event'])
  onMouseDown(event) {
    this.d3Service.applyMouseDownBehaviour();
    event.preventDefault();
  }
  svgMouseDown() {
    this.d3Service.applyMouseDownBehaviour();
  }
}
