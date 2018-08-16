import { Component, Input, Output, EventEmitter } from '@angular/core';
import { Node } from '../shared/node.model';

@Component({
  selector: '[nodeRect]',
  template: `
    <svg:g [attr.transform]="'translate(' + node.x + ',' + node.y + ')'" [attr.x]="node.x" [attr.y]="node.y">
      <svg:rect
          class="node"
          x="0" y="0"
          [attr.fill]="node.color"
          [attr.width]="width" [attr.height]="height"
          [attr.rx]="radius" [attr.ry]="radius">
      </svg:rect>
      <svg [attr.width]="width" [attr.height]="height">
        <svg:text
            class="node-type"
            x="50%"
            y="50%"
            [attr.font-size]="fontSize">
          {{node.type}}
        </svg:text>
      </svg>
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
          [attr.x]="width+7"
          [attr.y]="height/2"
          [attr.font-size]="fontSize">
        {{ node.name }}
      </svg:text>
    </svg:g>
  `,
  styleUrls: ['../shared/nodes.css']
})
export class NodeRectComponent {
  @Input('nodeRect') node: Node;
  @Output() clickRemove = new EventEmitter();

  width: number = 150;
  height: number = 25;
  radius: number = 4;
  deleteNodeX: number = this.width-5;
  deleteNodeY: number = -5;
  fontSize: number = 13;

  onClickRemove(node) {
    this.clickRemove.emit(node);
  }
}

