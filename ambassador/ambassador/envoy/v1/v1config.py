# Copyright 2018 Datawire. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License

from typing import Any, ClassVar, Dict, List, Optional, Tuple, Union

import json
import logging

from ...ir import IR

from .v1admin import V1Admin
from .v1listener import V1Listener
from .v1clustermanager import V1ClusterManager
from .v1tracing import V1Tracing

#############################################################################
## v1config.py -- the Envoy V1 configuration engine


class V1Config:
    def __init__(self, ir: IR) -> None:
        self.ir = ir

        # Toplevel stuff.
        self.admin: V1Admin = V1Admin.generate(self)

        # print("v1.admin %s" % self.admin)

        self.listeners: List[V1Listener] = V1Listener.generate(self)

        # print("v1.listeners %s" % self.listeners)

        self.clustermgr: V1ClusterManager = V1ClusterManager.generate(self)

        # print("v1.clustermgr %s" % self.clustermgr)
        self.tracing: Optional[V1Tracing] = V1Tracing.generate(self)


    def as_dict(self):
        d = {
            'admin': self.admin,
            'listeners': self.listeners,
            'cluster_manager': self.clustermgr,
        }

        if self.tracing:
            d['tracing'] = self.tracing

        return d

    def as_json(self):
        return json.dumps(self.as_dict(), sort_keys=True, indent=4)
