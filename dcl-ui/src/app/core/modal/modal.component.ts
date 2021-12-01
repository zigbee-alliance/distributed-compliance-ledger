import { Component, Input } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-modal',
  templateUrl: './modal.component.html',
  styleUrls: ['./modal.component.scss']
})
export class ModalWindowComponent {
  @Input() header;
  @Input() content;

  constructor(public activeModal: NgbActiveModal) {}
}
