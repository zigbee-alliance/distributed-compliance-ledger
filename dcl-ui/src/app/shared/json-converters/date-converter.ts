import { JsonConverter, JsonCustomConvert } from 'json2typescript';

@JsonConverter
export class DateConverter implements JsonCustomConvert<Date> {
  serialize(date: Date): any {
    return date ? date.toISOString() : date;
  }

  deserialize(date: any): Date {
    return new Date(date);
  }
}
