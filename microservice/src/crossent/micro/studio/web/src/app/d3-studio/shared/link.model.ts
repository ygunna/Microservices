import { Node } from './node.model';

export class Link {
  id: number;
  type: string;
  sNode: Node;
  tNode: Node;
  source: string;
  target: string;

  constructor(source, target) {
    this.sNode = source;
    this.tNode = target;
  }
}
