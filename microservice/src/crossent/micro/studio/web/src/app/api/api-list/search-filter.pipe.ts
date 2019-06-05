import { Pipe, PipeTransform } from '@angular/core';
import { MicroApi } from '../models/microapi.model'

@Pipe({
  name: 'searchApiFilter'
})
export class SearchApiFilterPipe implements PipeTransform {

  transform(microapis: MicroApi[], name:string, part:string, path:string): MicroApi[] {
    let ma: MicroApi[] = microapis;

    if(name && name != ""){
      ma = microapis.filter(microapi => microapi.name.indexOf(name) != -1  );
    }
    if(part && part != ""){
      ma = ma.filter(microapi => microapi.part == part );
    }
    if(path && path != ""){
      ma = ma.filter(microapi => microapi.path.indexOf(path) != -1  );
    }

    return ma;
  }

}
