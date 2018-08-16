import { Component, ElementRef, Input, Output, EventEmitter } from '@angular/core';
import { D3Service } from '../shared/d3-studio.service';
import { Link } from '../shared/link.model';

@Component({
  selector: '[linkPath]',
  template: `
    <svg:path
        class="link-path" [attr.id]="'link-path-'+link.id"
        [attr.d]="'M'+link.sNode.x+','+link.sNode.y+'L'+link.tNode.x+','+link.tNode.y" (click)="onClickPath(link)"
        style="marker-end: url(#end-arrow);"
    ></svg:path>
  `,
  styleUrls: ['./link-path.component.css']
})
export class LinkPathComponent  {
  @Input('linkPath') link: Link;
  @Output() clickPath = new EventEmitter();

  constructor(private d3Service: D3Service, private _element: ElementRef) { }

  onClickPath(link) {
    this.clickPath.emit(link);
    this.d3Service.clickPath(this._element.nativeElement, link);
  }
}
