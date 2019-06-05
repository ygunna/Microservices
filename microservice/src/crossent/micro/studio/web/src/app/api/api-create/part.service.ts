import { Injectable } from '@angular/core';
import { Part } from '../models/part.model';

@Injectable()
export class PartService {

  constructor() { }

  partsList(){
    let parts: Part[] = [
      {name:'', value:''},
      {name:'공공', value:'A'},
      {name:'교육', value:'B'},
      {name:'경제', value:'C'},
      {name:'광고', value:'D'},
      {name:'데이터', value:'E'},
      {name:'문화', value:'F'},
      {name:'미디어', value:'G'},
      {name:'비즈니스', value:'H'},
      {name:'소셜', value:'I'},
      {name:'스포츠', value:'J'},
      {name:'의료', value:'K'},
      {name:'여행', value:'L'},
      {name:'취미', value:'M'},
      {name:'기타', value:'Z'}
    ];
    return parts;
  }

}
