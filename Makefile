SRC = \
src/CoordCube.java \
src/CubieCube.java \
src/Search.java \
src/Server.java \
src/Tables.java \
src/Tools.java \
src/Util.java 

TESTSRC = \
test/Test.java \
test/TestServer.java \
test/TestTables.java

DIST = dist/src

DISTTEST = dist/test

ifndef probe
	probe = 0
endif

ifndef maxl
	maxl = 30
endif

ifndef ntest
	ntest = 1000
endif

.PHONY: build clean run testRnd testRndMP testRndStd testSel

build: $(DIST)

$(DIST): $(SRC) $(TESTSRC)
	@javac -d dist $(SRC) -Xlint:all
	@javac -d dist $(TESTSRC) -Xlint:all

run: $(DIST)
	@java -jar $(DIST)

serve: $(DIST)
	@nohup java -ea -cp dist:$(DIST) src.Server > dist/server.log &

testTables: $(DISTTEST)
	@java -ea -cp dist:$(DISTTEST) test.TestTables

testServer: $(DISTTEST)
	@java -ea -cp dist:$(DISTTEST) test.TestServer

testRnd: $(DISTTEST)
	@java -ea -cp dist:$(DIST) test.Test 40 $(ntest) $(maxl) 10000000 $(probe) 0

testRndMP: $(DISTTEST)
	@java -ea -cp dist:$(DIST) test.Test 72 $(ntest) $(maxl) 10000000 $(probe) 0

testRndStd: $(DISTTEST)
	@java -ea -cp dist:$(DIST) test.Test 40 $(ntest) 30 10000000 $(probe) 0 | grep AvgT
	@java -ea -cp dist:$(DIST) test.Test 40 $(ntest) 21 10000000 $(probe) 0 | grep AvgT
	@java -ea -cp dist:$(DIST) test.Test 40 $(ntest) 20 10000000 $(probe) 0 | grep AvgT

testSel: $(DISTTEST)
	@java -ea -cp dist:$(DIST) test.Test 24

$(DISTTEST): $(DIST) $(TESTSRC)
	@javac -d dist -cp dist:$(DIST) $(TESTSRC)

rebuild: clean build

clean:
	@rm -rf dist/*
