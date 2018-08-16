import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from "@angular/forms";
import { D3Service } from "./shared/d3-studio.service";
import { NodeCircleComponent } from "./node-circle/node-circle.component";
import { NodeRectComponent } from "./node-rect/node-rect.component";
import { LinkPathComponent } from "./link-path/link-path.component";
import { ZoomableDirective } from "./directives/zoomable.directive";
import { DraggableDirective } from "./directives/draggable.directive";
import { AppDraggableDirective } from "./directives/app-draggable.directive";
import { AppDroppableDirective } from "./directives/app-droppable.directive";

@NgModule({
  imports: [
    CommonModule, FormsModule
  ],
  declarations: [NodeCircleComponent, NodeRectComponent, LinkPathComponent, ZoomableDirective, DraggableDirective, AppDraggableDirective, AppDroppableDirective],
  providers: [D3Service],
  exports: [NodeCircleComponent, NodeRectComponent, LinkPathComponent, ZoomableDirective, DraggableDirective, AppDraggableDirective, AppDroppableDirective]
})
export class D3StudioModule { }
