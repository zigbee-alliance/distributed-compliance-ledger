/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Component, OnDestroy, OnInit } from '@angular/core';
import { BlockService } from '../../block.service';
import { BlockMeta } from '../../../../shared/models/block/block-meta';
import { Subscription, timer } from 'rxjs';

@Component({
  selector: 'app-block-list',
  templateUrl: './block-list.component.html',
  styleUrls: ['./block-list.component.scss']
})
export class BlockListComponent implements OnInit, OnDestroy {

  blockMetas: BlockMeta[] = [];

  collectionSize = 0;
  pageSize = 10;
  page = 1;

  timer: Subscription;

  constructor(private blockService: BlockService) {
  }

  ngOnInit() {
    this.timer = timer(0, 2000).subscribe(_ => {
      this.getBlockchainInfo();
    });
  }

  // TODO: Use async pipe
  getBlockchainInfo(): void {
    let minHeight = 0;
    let maxHeight = 0;

    // Not the first time
    if (this.collectionSize !== 0) {
      minHeight = this.collectionSize - this.page * this.pageSize + 1;
      maxHeight = minHeight + this.pageSize - 1;

      // Handle the start of block sequence
      if (minHeight < 1) {
        minHeight = 1;
      }
    }

    this.blockService.getBlockchainInfo(minHeight, maxHeight).subscribe(value => {
      this.blockMetas = value.blockMetas
        .sort((a, b) => b.header.height - a.header.height)
        .slice(0, this.pageSize);

      this.collectionSize = value.lastHeight;
    });
  }

  ngOnDestroy(): void {
    this.timer.unsubscribe();
  }
}
