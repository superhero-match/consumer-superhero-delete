/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package config

// Consumer holds the configuration values for the Kafka consumer.
type Consumer struct {
	Brokers []string `env:"KAFKA_BROKERS" yaml:"brokers" default:"[192.168.1.229:9092]"`
	Topic   string   `env:"KAFKA_DELETE_MUNICIPALITY_SUPERHERO_TOPIC" yaml:"topic" default:"delete.municipality.superhero"`
	GroupID string   `env:"KAFKA_SUPERHERO_DELETE_GROUP" yaml:"group_id" default:"consumer.superhero.delete.group"`
}
