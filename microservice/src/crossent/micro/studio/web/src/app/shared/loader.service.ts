import { Injectable, EventEmitter } from '@angular/core';

@Injectable()
export class LoaderService {
  changeActive: EventEmitter<boolean> = new EventEmitter<boolean>();
  count: number = 0;

  constructor() {}

  show() {
    this.count++;
    this.changeActive.emit(true);
  }

  hide() {
    this.count--;
    //console.log(this.count);
    if (this.count <= 0 ) {
      this.changeActive.emit(false);
    }
  }

  forceHide() {
    this.changeActive.emit(false);
  }

}

