import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ActivatedRoute } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { KeyInfo } from '../../../../shared/models/key/key-info';
import { KeyService } from '../../key.service';

@Component({
  selector: 'app-key-details',
  templateUrl: './key-details.component.html',
  styleUrls: ['./key-details.component.scss']
})
export class KeyDetailsComponent implements OnInit {

  item$: Observable<KeyInfo>;

  constructor(private route: ActivatedRoute,
              private keyService: KeyService) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      switchMap(params => this.keyService.getKeyInfo(params.get('name')))
    );
  }

}
