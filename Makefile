#
# Copyright 2011 Nathan Fiedler. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#
# Based on a sample by Dave Cheney via github.
#

include $(GOROOT)/src/Make.inc

CMDS=\
	src/cmd/mbench

PKGS=\
	src/pkg/sortingo

all: make

make: $(patsubst %, %.install, $(PKGS)) $(patsubst %, %.make, $(CMDS))
clean: $(patsubst %, %.clean, $(PKGS)) $(patsubst %, %.clean, $(CMDS))
nuke: $(patsubst %, %.nuke, $(PKGS)) $(patsubst %, %.nuke, $(CMDS))
test: $(patsubst %, %.test, $(PKGS)) $(patsubst %, %.test, $(CMDS))

%.install:
	$(MAKE) -C $* install

# In the case of multiple packages with a specific dependency order...
#package-2.install: package-1.install
#package-1.install package-2.install: package-3.install

%.make: %.install
	$(MAKE) -C $*

%.clean:
	$(MAKE) -C $* clean

%.nuke:
	$(MAKE) -C $* nuke

%.test:
	$(MAKE) -C $* test
