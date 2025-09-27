/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	"gorm.io/gorm"
	"seata.apache.org/seata-go-samples/quick_start/model"
	"time"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderDao(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (o *OrderService) Create(ctx context.Context, order model.Order) (int64, error) {
	now := time.Now().Unix()
	order.Ctime = now
	order.Utime = now
	err := o.db.WithContext(ctx).Create(order).Error
	if err != nil {
		return 0, err
	}
	return order.ID, nil
}

func (o *OrderService) Delete(ctx context.Context, id int64) error {
	return o.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", id).
		Update("is_deleted", "true").
		Error
}
