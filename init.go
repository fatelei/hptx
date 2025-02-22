/*
 * Copyright 2022 CECTC, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package hptx

import (
	"github.com/cectc/hptx/pkg/config"
	"github.com/cectc/hptx/pkg/core"
	"github.com/cectc/hptx/pkg/resource"
	"github.com/cectc/hptx/pkg/storage"
	"github.com/cectc/hptx/pkg/storage/etcd"
	"github.com/cectc/hptx/pkg/tcc"
)

func InitFromFile(path string) {
	conf := config.InitDistributedTransaction(path)
	resource.InitTCCBranchResource(tcc.GetResourceManager())
	driver := etcd.NewEtcdStore(conf.EtcdConfig)
	storage.InitStorageDriver(driver)
	core.InitDistributedTransactionManager(conf)
}

func InitWithConf(conf *config.DistributedTransaction) {
	resource.InitTCCBranchResource(tcc.GetResourceManager())
	driver := etcd.NewEtcdStore(conf.EtcdConfig)
	storage.InitStorageDriver(driver)
	core.InitDistributedTransactionManager(conf)
}
