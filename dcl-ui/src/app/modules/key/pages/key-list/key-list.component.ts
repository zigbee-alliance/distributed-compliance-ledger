import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { pluck, share } from 'rxjs/operators';
import { KeyService } from '../../key.service';
import { KeyInfo } from '../../../../shared/models/key/key-info';

@Component({
  selector: 'app-key-list',
  templateUrl: './key-list.component.html',
  styleUrls: ['./key-list.component.scss']
})
export class KeyListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<KeyInfo[]>;

  constructor(private keyService: KeyService) {
  }

  ngOnInit() {
    const source = this.keyService.getKeyInfos().pipe(
      share()
    );

    this.total$ = source.pipe(
      pluck('total')
    );

    this.items$ = source.pipe(
      pluck('items')
    );
  }

}
