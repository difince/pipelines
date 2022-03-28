# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import kfp_server_api
from kfp_server_api.api.run_service_api import RunServiceApi  # noqa: E501
from kfp_server_api.rest import ApiException


class TestRunServiceApi(unittest.TestCase):
    """RunServiceApi unit test stubs"""

    def setUp(self):
        self.api = kfp_server_api.api.run_service_api.RunServiceApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_archive_run(self):
        """Test case for archive_run

        Archives a run.  # noqa: E501
        """
        pass

    def test_create_run(self):
        """Test case for create_run

        Creates a new run.  # noqa: E501
        """
        pass

    def test_delete_run(self):
        """Test case for delete_run

        Deletes a run.  # noqa: E501
        """
        pass

    def test_get_run(self):
        """Test case for get_run

        Finds a specific run by ID.  # noqa: E501
        """
        pass

    def test_list_runs(self):
        """Test case for list_runs

        Finds all runs.  # noqa: E501
        """
        pass

    def test_list_runs_by_experiment(self):
        """Test case for list_runs_by_experiment

        """
        pass

    def test_read_artifact(self):
        """Test case for read_artifact

        Finds a run's artifact data.  # noqa: E501
        """
        pass

    def test_report_run_metrics(self):
        """Test case for report_run_metrics

        ReportRunMetrics reports metrics of a run. Each metric is reported in its own transaction, so this API accepts partial failures. Metric can be uniquely identified by (run_id, node_id, name). Duplicate reporting will be ignored by the API. First reporting wins.  # noqa: E501
        """
        pass

    def test_retry_run(self):
        """Test case for retry_run

        Re-initiates a failed or terminated run.  # noqa: E501
        """
        pass

    def test_terminate_run(self):
        """Test case for terminate_run

        Terminates an active run.  # noqa: E501
        """
        pass

    def test_unarchive_run(self):
        """Test case for unarchive_run

        Restores an archived run.  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
