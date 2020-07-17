PROJECT=delcol
TESTBIN=$(PROJECT)_test
TESTRUNNER=$(PROJECT)_testrunner.o
INCLUDEDIR=../include/common
CATCH=$(INCLUDEDIR)/catch.hpp

$(TESTBIN): $(TESTRUNNER) $(TESTBIN).cpp $(PROJECT).hpp
	g++ -I../include -Wall -o $@ $< $@.cpp

$(CATCH): $(INCLUDEDIR)
	[ -f $(CATCH) ] && touch $(CATCH) || curl -fLo $@ https://github.com/catchorg/Catch2/releases/download/v2.13.0/catch.hpp

$(INCLUDEDIR):
	mkdir -p $@

%.o: $(CATCH) %.cpp
	g++ -I../include -Wall -c $*.cpp

.PHONY: clean

clean:
	rm *.o $(TESTBIN)
