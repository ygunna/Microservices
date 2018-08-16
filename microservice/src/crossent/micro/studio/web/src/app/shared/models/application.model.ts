export class Application {
  metadata: {
    guid: string;
  };
  entity: {
    name: string;
    memory: number;
    disk_quota: number;
    buildpack: string;
    space_guid: string;
  };
}
