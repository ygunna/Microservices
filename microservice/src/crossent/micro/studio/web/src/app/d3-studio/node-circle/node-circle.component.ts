import { Component, Input, Output, EventEmitter } from '@angular/core';
import { Node } from '../shared/node.model';

@Component({
  selector: '[nodeCircle]',
  template: `
    <svg:g [attr.transform]="'translate(' + node.x + ',' + node.y + ')'" [attr.x]="node.x" [attr.y]="node.y">
      <svg:circle
          class="node"
          cx="0"
          cy="0"
          [attr.fill]="node.color"
          [attr.r]="radius">
      </svg:circle>
      <svg:text
          class="node-type"
          [attr.font-size]="fontSize">
        {{ node.type }}
      </svg:text>
      <svg id="close" [attr.x]="deleteNodeX" [attr.y]="deleteNodeY" class="delete-node invisible" (click)="onClickRemove(node)">
        <svg:line x1="1" y1="11" 
            x2="11" y2="1" 
            stroke="red" 
            stroke-width="2"/>
        <svg:line x1="1" y1="1" 
            x2="11" y2="11" 
            stroke="red" 
            stroke-width="2"/>
      </svg>
      <svg:text
          class="node-name"
          [attr.x]="radius+7"
          [attr.font-size]="fontSize">
        {{ node.name }}
      </svg:text>
    </svg:g>
  `,
  styleUrls: ['../shared/nodes.css']
})
export class NodeCircleComponent {
  @Input('nodeCircle') node: Node;
  @Output() clickRemove = new EventEmitter();

  radius: number = 25;
  deleteNodeX: number = this.radius;
  deleteNodeY: number = -(this.radius);
  fontSize: number = 13;

  onClickRemove(node) {
    this.clickRemove.emit(node);
  }
}
