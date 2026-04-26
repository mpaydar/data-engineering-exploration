
def generate_text_file():
    with open("logs.txt","r") as logs_file, open("report.txt","w") as report_file:
        report_file.write("Summary of logs\n")
        error_counter=0
        passed_counter=0
        for line in logs_file:
            if line.strip() == "Error":
                error_counter+=1
            else:
                passed_counter+=1
        report_file.write(f'Number of error counter {error_counter}\n')
        report_file.write(f'Number of passed counter {passed_counter}')
generate_text_file()
