class Solution {
public:
    double maxProbability(int n, vector<vector<int>>& edges, vector<double>& succProb, int start_node, int end_node) {
        vector<vector<pair<int, double>>> graph(n);
        vector<double> dist(n);
        for (int i = 0; i < edges.size(); i++) {
            graph[edges[i][0]].push_back({edges[i][1], succProb[i]});
            graph[edges[i][1]].push_back({edges[i][0], succProb[i]});
        }
        queue<int> q;
        q.push(start_node);
        dist[start_node] = 1;
        while(!(q.empty()) ) {
            int v = q.front(); q.pop();
            for (auto u : graph[v]) {
                if (dist[v] *  u.second > dist[u.first]) {
                    dist[u.first] = dist[v] * u.second;
                    q.push(u.first);
                }
            }
        }
        return dist[end_node];
    }
};