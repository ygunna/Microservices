import { Org } from './org.model';
import { Space } from './space.model';
import { Service } from './service.model';
import { Application } from './application.model';

export class Microservice {
  constructor(
    public id?: number,
    public org?: Org,
    public space?: Space,
    public name?: string,
    public version?: string,
    public description?: string,
    public visible?: string,
    public status?: string,
    public services?: {resources: Service[]},
    public apps?: {resources: Application[]}
  ) {}
}
