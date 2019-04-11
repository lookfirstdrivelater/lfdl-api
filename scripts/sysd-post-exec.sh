#!/usr/bin/env bash

sudo systemctl daemon-reload
sudo systemctl enable ldflapi.service
sudo  systemctl restart ldflapi.service
