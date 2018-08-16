import { Pipe, PipeTransform } from '@angular/core';
import { Micro } from './micro.model'

@Pipe({
  name: 'searchFilter'
})
export class SearchFilterPipe implements PipeTransform {

  transform(micros: Micro[], args:string): Micro[] {
    if(args == "") {
      return micros;
    }else{
      return micros.filter(micro => micro.name.indexOf(args) != -1  );
    }
  }

}
