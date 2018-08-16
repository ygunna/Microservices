import { Circle } from './circle.model'
import { Link } from './link.model'

export class Micro {
  constructor(
    public id: number,
    public orgName: string,
    public spaceName: string,
    public spaceGuid: string,
    public name: string,
    public version?: string,
    public description?: string,
    public app?: number,
    public service?: number,
    public status?: string,
    public url?: string,
    public swagger?: string,
  ) {}

  circles: Circle[];
  links: Link[];
}
