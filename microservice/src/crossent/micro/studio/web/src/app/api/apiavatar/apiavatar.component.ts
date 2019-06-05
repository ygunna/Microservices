import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: '.apiavatar',
  templateUrl: './apiavatar.component.html',
  styleUrls: ['./apiavatar.component.css']
})
export class ApiavatarComponent implements OnInit {
  @Input() image: string;
  @Input() x: number = 100;
  @Input() y: number = 100;

  constructor() { }

  ngOnInit() {

  }

}
