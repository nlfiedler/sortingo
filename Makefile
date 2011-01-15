#
# Copyright 2011 Nathan Fiedler. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#
# $Id: $
#

SUBDIRS = src/pkg/sortingo # src/cmd/sort src/cmd/bench

.PHONY: clean subdirs $(SUBDIRS)

subdirs: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@

clean:
# is there a better way?
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir clean; \
	done

test:
# is there a better way?
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir test; \
	done

# Declare dependency on package so commands are built last.
# src/cmd/sort: src/pkg/sortingo
# src/cmd/bench: src/pkg/sortingo

# TODO: add invocation of tago to build Emacs TAGS file
