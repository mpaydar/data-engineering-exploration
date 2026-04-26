
def generate_log_file():
    with open("logs.txt","w") as logs_file:
        for i in range(80):
            if i==2 or i==7:
                logs_file.write("Error\n")
            if i == 50:
                logs_file.write("Critical")
            else:
                logs_file.write("Passed\n")
generate_log_file()