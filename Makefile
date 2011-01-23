#
# Copyright 2011 Nathan Fiedler. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
#

SUBDIRS = src/pkg/sortingo src/cmd/mbench

.PHONY: all clean $(SUBDIRS)

all: install-pkg src/cmd/mbench

$(SUBDIRS):
	$(MAKE) -C $@

clean:
	for dir in $(SUBDIRS); do \
		$(MAKE) -C $$dir clean; \
	done

test: src/pkg/sortingo
	$(MAKE) -C src/pkg/sortingo test

install-pkg: src/pkg/sortingo
	$(MAKE) -C src/pkg/sortingo install

# TODO: add invocation of tago to build Emacs TAGS file
