#!/usr/bin/env bash

systemctl deamon-reload
systemctl enable ldflapi.service
systemctl restart ldflapi.service