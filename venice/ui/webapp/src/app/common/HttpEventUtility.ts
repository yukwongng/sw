import { Utility } from '@app/common/Utility';


/**
 * Create a new object for every watch endpoint in use.
 * Only checks by object name for uniqueness
 */
enum EventTypes {
  create = 'Created',
  update = 'Updated',
  delete = 'Deleted',
}

export class HttpEventUtility {
  private dataArray = [];
  // Maps object name to index in data array
  private dataMapping: { [objName: string]: number } = {};
  private filter: (object: any) => boolean;
  private objectConstructor: any;
  private isSingleton: boolean;

  /**
   * @param objectConstructor   Constructor that will be called on all
   *                            incoming data if given
   *
   * @param isSingleton         Whether the events are for a singleton object
   *                            Supports object being renamed if its a singleton
   *
   * @param filter              If the filter returns false for an object,
   *                            it won't be added to the array
   */
  constructor(objectConstructor: any = null, isSingleton: boolean = false, filter: (object: any) => boolean = null) {
    this.objectConstructor = objectConstructor;
    this.isSingleton = isSingleton;
    this.filter = filter;
  }

  /**
   * Can be used by other components as an efficient way to
   * check if the array has changed
   */
  public static trackBy(index: number, item: any): string {
    return item.meta.name + ' - ' + item.meta['mod-time'];
  }

  public processEvents(eventChunk) {
    if (eventChunk.result == null) {
      console.log('event chunk was blank', eventChunk);
      return;
    }
    const events = eventChunk.result.Events;
    events.forEach(event => {
      let obj;
      if (this.objectConstructor != null) {
        obj = new this.objectConstructor(event.Object);
      } else {
        obj = event.Object;
      }
      const objName = obj.meta.name;
      if (this.filter != null && !this.filter(obj)) {
        return;
      }
      let index;
      switch (event.Type) {
        case EventTypes.create:
          index = this.dataArray.length;
          this.dataArray.push(obj);
          this.dataMapping[objName] = index;
          break;
        case EventTypes.delete:
          index = this.dataMapping[objName];
          this.dataArray.splice(index, 1);
          delete this.dataMapping[objName];
          // Decrement index of every element after
          // the one we removed
          for (const key in this.dataMapping) {
            if (this.dataMapping.hasOwnProperty(key)) {
              const value = this.dataMapping[key];
              if (value > index) {
                this.dataMapping[key] = value - 1;
              }
            }
          }
          break;
        case EventTypes.update:
          if (this.isSingleton && this.dataArray.length > 0) {
            // Can only be one object, so all updates are happening
            // to the one object we have
            index = 0;
          } else {
            index = this.dataMapping[objName];
          }
          if (index != null) {
            this.dataArray[index] = obj;
          } else {
            console.log('Update event but object was not found');
          }
          break;
        default:
          break;
      }
    });

    return this.dataArray;
  }

  public hasItem(objName: string) {
    return this.dataMapping[objName];
  }

  public get array(): ReadonlyArray<any> {
    return this.dataArray;
  }

}
