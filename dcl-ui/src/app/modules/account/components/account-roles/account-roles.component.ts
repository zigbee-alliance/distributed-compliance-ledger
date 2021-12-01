import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-account-roles',
  templateUrl: './account-roles.component.html',
  styleUrls: ['./account-roles.component.scss']
})
export class AccountRolesComponent implements OnInit {

  @Input() roles: string[];

  constructor() {
  }

  ngOnInit(): void {
  }

}
