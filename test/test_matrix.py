# Test matrix file - load in the matrix jsons and send a request to 
# the matrix adder
import json
import requests

def test_matrix_multiplier():

    # import the matrix files
    with open("test/matrix_a.json") as f:
        matrix_a = json.load(f)
    with open("test/matrix_b.json") as f:
        matrix_b = json.load(f)

    # set up the request
    j_request = []
    j_request.append(matrix_a)
    j_request.append(matrix_b)
    
    # send matrixes
    r = requests.get('http://127.0.0.1:8099/multi', data=json.dumps(j_request))
    
    # assert the response
    print(f"The response is {r.text}")
    assert(r.status_code == 200)
    assert(int(r.text) == 96)