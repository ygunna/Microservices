import { Pipe, PipeTransform } from '@angular/core';
import { MicroApi } from '../models/microapi.model'
import { Part } from '../models/part.model';
import { PartService } from '../api-create/part.service';

@Pipe({
  name: 'codeFilter'
})
export class CodeFilterPipe implements PipeTransform {

  constructor(private partService: PartService) { }

  transform(part: MicroApi): string {
    let parts = this.partService.partsList().filter(d => d.value == part);
    if (parts.length > 0) return parts[0].name;
    else return '';
  }

}
