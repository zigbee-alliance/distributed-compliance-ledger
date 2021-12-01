import { JsonConverter, JsonCustomConvert } from 'json2typescript';

// Amino codec expects quoted values for javascript numeric type
@JsonConverter
export class AminoNumberConverter implements JsonCustomConvert<number> {

  serialize(value: number): any {
    return value.toString();
  }

  deserialize(value: any): number {
    return value ? parseInt(value, 10) : 0;
  }
}
