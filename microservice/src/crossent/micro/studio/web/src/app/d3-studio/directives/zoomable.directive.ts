import { Directive, Input, ElementRef, OnInit } from '@angular/core';
import { D3Service } from '../shared/d3-studio.service';

@Directive({
  selector: '[zoomableOf]'
})
export class ZoomableDirective implements OnInit {
  @Input('zoomableOf') zoomableOf: ElementRef;

  constructor(private d3Service: D3Service, private _element: ElementRef) {}

  ngOnInit() {
    this.d3Service.applyZoomableBehaviour(this.zoomableOf, this._element.nativeElement);
  }
  zoomClick(direction) {
    this.d3Service.zoomClick(direction);
  }
  nodeSimulation(width, height, nodes, links) {
    this.d3Service.nodeSimulation(width, height, nodes, links);
  }
}
