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
      <svg id="info" [attr.x]="infoNodeX" [attr.y]="infoNodeY" class="delete-node invisible" (click)="onClickInfo(node)">
        <svg:path style="fill:#03A9F4;" d="M7,0C3.2,0,0,3.2,0,7s3.2,7,7,7s7-3.1,7-7S10.9,0,7,0z M7,3.5c0.5,0,0.9,0.3,0.9,0.9c0,0.5-0.3,0.9-0.9,0.9
          S6.2,4.9,6.2,4.4C6.2,3.8,6.5,3.5,7,3.5z M8.3,10.5H5.7c-0.2,0-0.5-0.2-0.5-0.5s0.2-0.5,0.5-0.5h0.5V7H5.7C5.5,7,5.3,6.8,5.3,6.5
          c0-0.2,0.2-0.5,0.5-0.5h1.8c0.2,0,0.5,0.2,0.5,0.5v3h0.5c0.2,0,0.5,0.2,0.5,0.5C8.8,10.3,8.5,10.5,8.3,10.5z"/>
      </svg>
      <svg id="close" [attr.x]="deleteNodeX" [attr.y]="deleteNodeY" class="delete-node invisible" (click)="onClickRemove(node)">
        <svg:path style="fill:#F85359;" d="M7,0C3.1,0,0,3.1,0,7s3.1,7,7,7c3.9,0,7-3.1,7-7S10.9,0,7,0z M9.5,8.9c0.2,0.2,0.2,0.4,0,0.6
				C9.3,9.6,9,9.6,8.9,9.5L7,7.6L5.1,9.5C5,9.7,4.7,9.7,4.5,9.5C4.3,9.3,4.3,9,4.5,8.9L6.4,7L4.5,5.1C4.4,5,4.4,4.7,4.5,4.5
				c0.2-0.2,0.4-0.2,0.6,0L7,6.4l1.9-1.9c0.2-0.2,0.5-0.2,0.6,0c0.2,0.2,0.2,0.5,0,0.6L7.6,7L9.5,8.9z"/>
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
  @Output() clickInfo = new EventEmitter();
  @Output() clickRemove = new EventEmitter();

  radius: number = 25;
  deleteNodeX: number = this.radius;
  deleteNodeY: number = -(this.radius);
  infoNodeX: number = this.deleteNodeX;
  infoNodeY: number = this.deleteNodeY-15;
  fontSize: number = 13;

  onClickInfo(node) {
    this.clickInfo.emit(node);
  }
  onClickRemove(node) {
    this.clickRemove.emit(node);
  }
}
