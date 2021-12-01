import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { pluck, share } from 'rxjs/operators';
import { Validator } from 'src/app/shared/models/validator/validator';
import { ValidatorService } from '../../validator.service';

@Component({
  selector: 'app-validator-list',
  templateUrl: './validator-list.component.html',
  styleUrls: ['./validator-list.component.scss']
})
export class ValidatorListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<Validator[]>;

  constructor(public validatorService: ValidatorService) {
  }

  ngOnInit() {
    this.getValidatorHeaders();
  }

  getValidatorHeaders(): void {
    const source = this.validatorService.getAllValidators().pipe(share());

    this.total$ = source.pipe(pluck('total'));
    this.items$ = source.pipe(pluck('items'));
  }

}
