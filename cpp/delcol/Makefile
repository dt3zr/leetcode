PROJECT=$(shell basename $$(pwd))
BUILDDIR=build
TESTBINNAME=$(PROJECT)_test
TESTBIN=$(BUILDDIR)/$(TESTBINNAME)
TESTRUNNER=$(BUILDDIR)/$(PROJECT)_testrunner.o
INCLUDEDIR=../include/common
CATCH=$(INCLUDEDIR)/catch.hpp

$(TESTBIN): $(TESTRUNNER) $(TESTBINNAME).cpp $(PROJECT).hpp
	g++ -I../include -Wall -o $@ $< $(TESTBINNAME).cpp

$(INCLUDEDIR):
	mkdir -p $@

$(CATCH): $(INCLUDEDIR)
	[ -f $(CATCH) ] && touch $(CATCH) || curl -fLo $@ https://github.com/catchorg/Catch2/releases/download/v2.13.0/catch.hpp

$(BUILDDIR)/%.o: $(CATCH) $(BUILDDIR) %.cpp
	g++ -I../include -Wall -o $(TESTRUNNER) -c $*.cpp

$(BUILDDIR):
	mkdir build

.PHONY: clean

clean:
	rm -rf build
