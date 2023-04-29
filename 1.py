import sys

filename = sys.argv[1]
f = open(filename, "a")
f.write("appended from python")
f.close()