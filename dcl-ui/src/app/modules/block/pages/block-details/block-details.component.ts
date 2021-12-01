import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { TxService } from '../../../tx/tx.service';
import { pluck, share, switchMap } from 'rxjs/operators';
import { BlockService } from '../../block.service';
import { ResultBlock } from '../../../../shared/models/block/result-block';
import { DecodeTxsResponse } from '../../../../shared/models/tx/decode-txs-response';

@Component({
  selector: 'app-block-details',
  templateUrl: './block-details.component.html',
  styleUrls: ['./block-details.component.scss']
})
export class BlockDetailsComponent implements OnInit {
  block$: Observable<ResultBlock>;
  txs$: Observable<any[]>;

  constructor(private route: ActivatedRoute,
              private blockService: BlockService,
              private txService: TxService) {
  }

  ngOnInit() {
    this.block$ = this.route.paramMap.pipe(
      switchMap(params => this.blockService.getBlock(+params.get('height'))),
      share()
    );

    this.txs$ = this.block$.pipe(
      switchMap(resultBlock => this.txService.decodeTx(resultBlock.block.data)),
      pluck('txs')
    );
  }
}
